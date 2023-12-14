import { NextResponse, NextRequest } from "next/server";

export async function middleware(request: NextRequest) {
  const cookie = request.cookies.get("session_token");
  if (!cookie) {
    return NextResponse.redirect(new URL("/signin", request.url));
  }
  return NextResponse.next();
}
export const config = {
  matcher: "/user/:path*",
};
