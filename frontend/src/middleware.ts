import { NextResponse, NextRequest } from "next/server";
import axios from "@/api/axiosConfig";
import { useAuth } from "@/app/userContext";
// import type { NextRequest } from "next/server";

export async function middleware(request: NextRequest, response: NextResponse) {
  const cookie = request.cookies.get("session_token");
  console.log("cookie", cookie);
  if (!cookie) {
    return NextResponse.redirect(new URL("/signin", request.url));
  }
  return NextResponse.next();
}
export const config = {
  matcher: "/user/:path*",
};
