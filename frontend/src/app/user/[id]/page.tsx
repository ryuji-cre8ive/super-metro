"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { jwtDecode } from "jwt-decode";

export default function Page({ params }: { params: { id: string } }) {
  const router = useRouter();
  return (
    <div>
      <p>youre id is : {params.id}</p>
      <p>youre token is: {window.localStorage.getItem("session_token")}</p>
      <button onClick={() => router.push(`/user/${params.id}/signout`)}>
        SignOut
      </button>
    </div>
  );
}
