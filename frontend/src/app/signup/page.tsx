"use client";

import React from "react";
import LoginForm from "@/components/login";

const onSubmit = (email: string, password: string) => {
  console.log(`Email: ${email}, Password: ${password}`);
};

export default function signup() {
  return (
    <main>
      <h1>This is a sign up form</h1>
      <LoginForm onSubmit={onSubmit} />
    </main>
  );
}
