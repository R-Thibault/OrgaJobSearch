// /src/app/context/UserContext.tsx
"use client";

import { createContext, useContext, useEffect, useState } from "react";
import axios from "axios";
import { UserType } from "@/types/userTypes";

interface UserContextType {
  user: UserType | null;
  loading: boolean;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<UserType | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchUserRole = async () => {
      try {
        const response = await axios.get("http://localhost:8080/me", {
          withCredentials: true,
        });
        if (response.status === 200) {
          const userProfile = response.data;
          setUser(userProfile);
        }
      } catch (error) {
        console.error("Error:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchUserRole();
  }, []);

  return (
    <UserContext.Provider value={{ loading, user }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error("useUser must be used within a UserProvider");
  }
  return context;
};
