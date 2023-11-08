"use client";
import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from "react";
import { User } from "@/app/models/user";
import { jwtDecode } from "jwt-decode";
// Contextを作成
const AuthContext = createContext({
  user: {} as User | null | undefined,
  login: (user: User) => {},
  logout: () => {},
  topUp: (user: User, amount: number) => {},
});

// AuthProviderコンポーネントを作成
export const AuthProvider = ({ children }: { children: ReactNode }) => {
  useEffect(() => {
    const token = window.localStorage.getItem("session_token");
    if (!token) {
      return logout();
    }
    const decodedToken = jwtDecode(token) as User;
    if (!decodedToken) {
      return;
    }
    const userInfo: User = {
      id: decodedToken.id,
      password: "",
      email: decodedToken.email,
      userName: decodedToken.userName,
      valance: decodedToken.valance,
      sessionToken: "",
    };

    setUser(userInfo);
  }, []);
  const [user, setUser] = useState<User | null>();

  const login = (user: User) => {
    setUser(user);
  };

  const logout = () => {
    setUser(null);
  };

  const topUp = (user: User, amount: number) => {
    setUser({ ...user, valance: user.valance + amount });
  };

  return (
    <AuthContext.Provider value={{ user, login, logout, topUp }}>
      {children}
    </AuthContext.Provider>
  );
};

// useContextのカスタムフックを作成
export const useAuth = () => {
  return useContext(AuthContext);
};
