"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { jwtDecode } from "jwt-decode";
import axios from "@/api/axiosConfig";

export default function Page({ params }: { params: { id: string } }) {
  const router = useRouter();
  const onSignOut = () => {
    if (!confirm("are you sure you want to sign out?")) {
      return;
    }
    window.localStorage.removeItem("session_token");
    axios.post("/logout");
    router.push("/");
  };

  useEffect(() => {
    const token = window.localStorage.getItem("session_token");
    if (!token) {
      router.push("/");
      return;
    }

    try {
      const decoded = jwtDecode(token);
      console.log("decoded", decoded);
      if (!decoded) {
        router.push("/");
      }
    } catch (err) {
      router.push("/");
    }
  }, []);

  return (
    <div>
      <p>youre id is : {params.id}</p>
      <p>youre token is: {window.localStorage.getItem("session_token")}</p>
      <button onClick={onSignOut}>SignOut</button>
    </div>
  );
}
