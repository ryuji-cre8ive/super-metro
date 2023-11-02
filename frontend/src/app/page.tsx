"use client";
import React from "react";
import LoginForm from "@/components/login";
import Link from "next/link";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";
import "./page.module.css";

export default function Home() {
  const router = useRouter();

  const onSubmit = async (email: string, password: string) => {
    console.log(`Email: ${email}, Password: ${password}`);
    const res: AxiosResponse = await axios.post("/login", { email, password });
    console.log("res", res);
    if (res.status === 200) {
      localStorage.setItem("session_token", res.data.sessionToken);
      router.push(`/user/${res.data.id}`);
    }
  };

  return (
    <main>
      <h1>Hello super-metro</h1>
      <h2>This is a login form </h2>
      <LoginForm onSubmit={onSubmit} />
      <Link href={"/signup"}>still not have account??</Link>
    </main>
  );
}
