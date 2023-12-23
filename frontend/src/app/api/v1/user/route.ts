import { NextRequest, NextResponse } from "next/server";
import { jwtDecode } from "jwt-decode";
import { User } from "@/app/models/user";

export async function GET(request: NextRequest) {
  if (!request.url) {
    return new NextResponse(null, { status: 401 });
  }
  try {
    const cookie = request.cookies.get("session_token");
    if (!cookie) {
      return new NextResponse(null, { status: 401 });
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
    return NextResponse.error();
  }
}
