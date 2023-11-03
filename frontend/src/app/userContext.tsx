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
});

// AuthProviderコンポーネントを作成
export const AuthProvider = ({ children }: { children: ReactNode }) => {
  useEffect(() => {
    const token = localStorage.getItem("session_token");
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

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// useContextのカスタムフックを作成
export const useAuth = () => {
  return useContext(AuthContext);
};
