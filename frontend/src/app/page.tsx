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

  return (
    <main>
      <h1>Hello super-metro</h1>
      <h2>This is top page </h2>
    </main>
  );
}
