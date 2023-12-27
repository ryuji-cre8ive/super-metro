"use client";
import React, { useState } from "react";
import LoginForm from "@/components/login";
import Link from "next/link";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";
import { useAuth } from "@/app/userContext";
import { User } from "@/app/models/user";
import CustomSnackBar from "@/components/SnackBar";

export default function Home() {
  const router = useRouter();
  const [open, setOpen] = useState<boolean>(false);
  const [message, setMessage] = useState<string>("");
  const { user, login, logout } = useAuth();

  const onSubmit = async (email: string, password: string) => {
    try {
      const res: AxiosResponse = await axios.post("/login", {
        email,
        password,
      });
      if (res.status === 200 && res.data.status !== 401) {
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
      } else {
        setOpen(true);
        setMessage(res.data.message);
      }
    } catch (err) {
      console.log("login failed", err);
    }
  };

  return (
    <main>
      <LoginForm onSubmit={onSubmit} />
      <CustomSnackBar
        open={open}
        setOpen={setOpen}
        severity="error"
        text={message}
      />
    </main>
  );
}
