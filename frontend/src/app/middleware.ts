import { NextResponse } from "next/server";
import { useAuth } from "@/app/userContext";
// import type { NextRequest } from "next/server";

export function middleware() {
  const token = localStorage.getItem("session_token");
  if (!token) {
    return NextResponse.redirect("/signin");
  }
}
export const config = {
  matcher: "/user/*",
};
