"use client";
import React from "react";
import Signup from "@/components/signup";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";

export default function SignupPage() {
  const router = useRouter();

  const onSubmit = (email: string, password: string, userName: string) => {
    const params = {
      email,
      password,
      userName,
    };

    axios.post("/signup", params).then((res: AxiosResponse) => {
      alert("successfully created !!!");
      if (res.status === 200) {
        return router.push("/");
      }
      return alert("something went wrong");
    });
  };

  return (
    <main>
      <h1>This is a sign up form</h1>
      <Signup onSubmit={onSubmit} />
    </main>
  );
}
