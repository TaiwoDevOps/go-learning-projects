"use client";

import ChatBody from "@/components/chat_body";
import { useEffect, useContext, useState, useRef } from "react";
import { WebSocketContext } from "@/modules/websocket_provider";
import { AuthContext } from "@/modules/auth_provider";
import { useRouter } from "next/navigation";
import { API_URL } from "@/constants";
import autosize from "autosize";

export type Message = {
  content: string;
  client_id: string;
  username: string;
  room_id: string;
  type: "self" | "other";
};
const Index = () => {
  const [messages, setMessages] = useState<Array<Message>>([]);
  const textarea = useRef<HTMLTextAreaElement>(null);
  const { conn } = useContext(WebSocketContext);
  const { user } = useContext(AuthContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);

  const router = useRouter();

  useEffect(() => {
    if (conn === null) {
      router.push("/");
      return;
    }

    const roomID = conn.url.split("/")[5];
    console.log(conn.url);
    console.log("******==========>>>><<<<<==========******");
    async function getUsers() {
      try {
        const res = await fetch(`${API_URL}/ws/clients/${roomID}`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });

        const data = await res.json();

        setUsers(data);
      } catch (error) {
        console.error(error);
        console.log("******==========>>>><<<<<==========******");
      }
    }
    getUsers();
  }, [conn, router]);

  useEffect(() => {
    if (textarea.current) {
      autosize(textarea.current);
    }

    if (conn === null) {
      router.push("/");
      return;
    }

    conn.onmessage = (e) => {
      const m: Message = JSON.parse(e.data);
      if (m.content.includes("joined")) {
        setUsers([...users, { username: m.username }]);
      }

      if (m.content.includes("left")) {
        const deleteUser = users.filter((u) => u.username != m.username);
        setUsers([...deleteUser]);
        setMessages([...messages, m]);
        return;
      }

      if (m.client_id === user.id) {
        m.type = "self";
      } else {
        m.type = "other";
      }
      setMessages([...messages, m]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textarea, messages, conn, users, router, user.id]);

  const sendMessage = () => {
    console.log(`Sending message ${textarea.current?.value}`);
    console.log("******==========>>>><<<<<==========******");

    if (!textarea.current?.value) return;
    if (conn === null) {
      router.push("/");
      return;
    }

    conn.send(textarea.current.value);
    textarea.current.value = "";
    autosize.update(textarea.current);
    textarea.current.focus();
  };

  return (
    <>
      <div className="flex flex-col w-full">
        <div className="p-4 md:mx-6 mb-14">
          <ChatBody data={messages} />
        </div>
        <div className="fixed bottom-0 mt-4 w-full">
          <div className="flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md">
            <div className="flex w-full mr-4 rounded-md border border-blue">
              <textarea
                ref={textarea}
                placeholder="Type your message here..."
                className="w-full h-10 p-2 rounded-md focus:outline-none"
                style={{ resize: "none" }}
              />
            </div>
            <div className="flex items-center">
              <button
                onClick={sendMessage}
                // disabled={!textarea.current?.value}
                className="p-2 rounded-mb bg-blue text-white"
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Index;
