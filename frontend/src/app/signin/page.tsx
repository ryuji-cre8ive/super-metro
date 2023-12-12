"use client";
import React from "react";
import LoginForm from "@/components/login";
import Link from "next/link";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";
import { useAuth } from "@/app/userContext";
import { User } from "@/app/models/user";

export default function Home() {
  const router = useRouter();
  const { user, login, logout } = useAuth();

  const onSubmit = async (email: string, password: string) => {
    const res: AxiosResponse = await axios.post("/login", { email, password });
    if (res.status === 200) {
      // document.cookie = `session_token=${res.data.sessionToken}; path=/;`;
      const userInfo: User = {
        id: res.data.id,
        userName: res.data.userName,
        email: res.data.email,
        password: res.data.password,
        valance: res.data.valance,
        sessionToken: res.data.sessionToken,
      };

      login(userInfo);
      router.push(`/user/${res.data.id}`);
    }
  };

  return (
    <main>
      <h2>This is a login form </h2>
      <LoginForm onSubmit={onSubmit} />
    </main>
  );
}
