"use client";
import React from "react";
import LoginForm from "@/components/login";
import Link from "next/link";

const onSubmit = (email: string, password: string) => {
  console.log(`Email: ${email}, Password: ${password}`);
};

export default function Home() {
  return (
    <main>
      <h1>Hello super-suica</h1>

      <h2>This is a login form </h2>
      <LoginForm onSubmit={onSubmit} />
      <Link href={"/signup"}>still not have account??</Link>
    </main>
  );
}
