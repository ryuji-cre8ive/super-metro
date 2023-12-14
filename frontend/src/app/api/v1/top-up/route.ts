import { NextRequest, NextResponse } from "next/server";

import axios from "axios";

export async function POST(request: NextRequest) {
  if (!request.url) {
    return NextResponse.error();
  }
  let url = request.url.replace("localhost:3000", "localhost:8080");
  let body;
  try {
    body = await request.json();
  } catch (error) {
    console.error(error);
    return NextResponse.error();
  }
  const config = {
    method: request.method.toLowerCase(),
    url: url,
    headers: Object.fromEntries(request.headers.entries()),
    data: body,
  };
  try {
    const res = await axios(config);
    const bffRes = NextResponse.json(res.data);
    bffRes.cookies.set("session_token", res.data.sessionToken, {
      httpOnly: true,
      path: "/",
      expires: new Date(Date.now() + 1000 * 60 * 30),
      sameSite: "lax",
    });
    return bffRes;
  } catch (error) {
    return NextResponse.error();
  }
}
