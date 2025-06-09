"use client";

import { useState, useEffect, useContext } from "react";
import { API_URL,WEBSOCKET_URL } from "../constants";
import { v4 as uuid4 } from "uuid";
import { AuthContext } from "@/modules/auth_provider";
import {WebSocketContext} from "@/modules/websocket_provider";
import {useRouter } from 'next/navigation' 


const Index = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([]);
  const [roomName, setRoomName] = useState<string>("");
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebSocketContext);


  const router = useRouter()

  const getRooms = async () => {
    try {
      setRoomName("");
      const res = await fetch(`${API_URL}/ws/rooms`, {
        method: "GET",
      });

      if (res.ok) {
        setRooms(await res.json());
      }
    } catch (err) {
      console.error("Error creating room:", err);
    }
  };

  useEffect(() => {
    getRooms();
  }, []);

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      setRoomName("");
      const body = JSON.stringify({
        id: uuid4(),
        name: roomName,
      });
          console.log('submitHandler called', body);
      console.log('★★★★ End create room submitHandler inner soldiers ★★★★')
      const res = await fetch(`${API_URL}/ws/create-room`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body:  body,
      });

      if (!res.ok) {
        throw new Error("Failed to create room");
      } else {
        getRooms();
      }
    } catch (err) {
      console.error("Error creating room:", err);
    }
  };

  const joinRoom = async (roomId: string) => {
  const ws = new WebSocket(`${WEBSOCKET_URL}/ws/join-room/${roomId}?userId=${user.id}&userName=${user.userName}`);
  if (ws.OPEN){
    setConn(ws);
    router.push(`/chat`);
    return
  }
  };

  return (
    <>
      <div className="my-8 px-4 md:mx-32 w-full h-full">
        <div className="flex justify-center mt-3 p-5">
          <input
            placeholder="Enter room name"
            type="text"
            className="border border-grey p-2 rounded-md focus:outline-none focus:border-blue"
            value={roomName}
            onChange={(e) => setRoomName(e.target.value)}
          />

          <button
            className="bg-blue border text-white rounded-md p-2 md:ml-4"
            onClick={submitHandler}
          > 
            Create Room
          </button>
        </div>
        <div className='mt-6'>
          <div className='font-bold'>Available Rooms</div>
          <div className='grid grid-cols-1 md:grid-cols-3 gap-4 mt-6'>  
            {rooms.map((room, index) => (
              <div
                key={index}
                className='border border-blue p-4 flex items-center rounded-md w-full'
              >
                <div className='w-full'>
                  <div className='text-sm'>Room info:</div>
                  <div className='text-blue font-bold text-lg'>{room.name}</div>
                </div>
                <div className=''>
                  <button 
                  onClick={() => {
                   joinRoom(room.id);
                  }
                  }
                  className='px-4 text-white bg-blue rounded-md'>Join</button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

// time stamp : 21:30 'commence sign up screen dev" starting with dashboard screen
export default Index;
