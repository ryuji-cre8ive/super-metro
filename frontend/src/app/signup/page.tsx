"use client";
import React, { useState } from "react";
import Signup from "@/components/signup";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";
import CustomSnackBar from "@/components/SnackBar";

export default function SignupPage() {
  const [open, setOpen] = useState<boolean>(false);
  const [message, setMessage] = useState<string>("");
  const router = useRouter();

  const onSubmit = (email: string, password: string, userName: string) => {
    const params = {
      email,
      password,
      userName,
    };

    axios.post("/signup", params).then((res: AxiosResponse) => {
      if (res.status === 200 && res.data.status !== 400) {
        alert("user was created successfully");
        return router.push("/");
      }
      setOpen(true);
      setMessage(res.data.message);
      console.log(res.data);
      return;
    });
  };

  return (
    <main>
      <h1>This is a sign up form</h1>
      <Signup onSubmit={onSubmit} />
      <CustomSnackBar
        open={open}
        setOpen={setOpen}
        severity="error"
        text={message}
      />
    </main>
  );
}
