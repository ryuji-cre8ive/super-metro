"use client";
import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from "react";
import { useCookies } from "react-cookie";
import { User } from "@/app/models/user";
import { jwtDecode } from "jwt-decode";
import axios from "@/api/axiosConfig";
// Contextを作成
const AuthContext = createContext({
  user: {} as User | null | undefined,
  login: (user: User) => {},
  logout: () => {},
  topUp: (user: User, amount: number) => {},
});

// AuthProviderコンポーネントを作成
export const AuthProvider = ({ children }: { children: ReactNode }) => {
  // const [cookies] = useCookies<string>(["session_token"]);
  // console.log("token", cookies.session_token);
  // const token = cookies.session_token;
  const getUserInfo = async () => {
    const res = await axios.get("/user");
    console.log(res.data);
    return res.data;
  };

  useEffect(() => {
    const fetchData = async () => {
      const data = await getUserInfo();
      if (!data) {
        return logout();
      }
      const userInfo: User = {
        id: data.id,
        password: "",
        email: data.email,
        userName: data.userName,
        valance: data.valance,
        sessionToken: "",
      };
      console.log(userInfo);

      setUser(userInfo);
    };

    fetchData();
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
