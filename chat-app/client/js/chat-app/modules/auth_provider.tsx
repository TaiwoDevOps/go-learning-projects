'use client';

import { useState, createContext, useEffect } from "react";
import { useRouter } from "next/navigation";

export type UserInfo = {
  userName: string;
  id: string;
};

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: (auth: boolean) => void;
  user: UserInfo;
  setUser: (user: UserInfo) => void;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: {
    userName: "",
    id: "",
  },
  setUser: () => {},
});

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUser] = useState<UserInfo>({
    userName: "",
    id: "",
  });

  const router = useRouter();

  useEffect(() => {
    const userInfo = localStorage.getItem("user_info");
    if (!userInfo) {
      if (window.location.pathname != "/signup") {
        router.push("/login");
        return;
      }
    } else {
      const user: UserInfo = JSON.parse(userInfo);
      if (user) {
        setUser({
          userName: user.userName,
          id: user.id,
        });
      }
      setAuthenticated(true);
    }
  }, [authenticated, router]);

  return (
    <AuthContext.Provider
      value={{
        authenticated: authenticated,
        setAuthenticated: setAuthenticated,
        user: user,
        setUser: setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContextProvider;
