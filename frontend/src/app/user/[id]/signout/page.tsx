"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import axios from "@/api/axiosConfig";
import Button from "@mui/material/Button";
import { useAuth } from "@/app/userContext";

export default function SignOutPage() {
  const { logout } = useAuth();
  const router = useRouter();
  const onSignOut = async () => {
    if (!confirm("are you sure you want to sign out?")) {
      return;
    }
    window.localStorage.removeItem("session_token");
    logout();
    await axios.post("/logout");
    router.push("/");
  };
  return (
    <>
      <Button onClick={onSignOut}>Logout</Button>
    </>
  );
}
