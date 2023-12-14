import { NextRequest, NextResponse } from "next/server";
import { jwtDecode } from "jwt-decode";
import { User } from "@/app/models/user";

export async function GET(request: NextRequest) {
  if (!request.url) {
    return NextResponse.error();
  }
  try {
    const cookie = request.cookies.get("session_token");
    if (!cookie) {
      return NextResponse.error();
    }

    const decodedToken = jwtDecode(cookie.value) as User;
    if (!decodedToken) {
      return NextResponse.error();
    }
    const userInfo: User = {
      id: decodedToken.id,
      password: "",
      email: decodedToken.email,
      userName: decodedToken.userName,
      valance: decodedToken.valance,
      sessionToken: "",
    };
    return NextResponse.json(userInfo);
  } catch (error) {
    console.error(error);
    return NextResponse.error();
  }
}
