import { NextRequest, NextResponse } from "next/server";

import axios from "axios";

export async function GET(request: NextRequest) {
  if (!request.url) {
    return NextResponse.error();
  }
  let url = request.url.replace("localhost:3000", "localhost:8080");
  const config = {
    method: request.method.toLowerCase(),
    url: url,
    headers: Object.fromEntries(request.headers.entries()),
    data: {},
  };
  try {
    const res = await axios(config);
    return NextResponse.json(res.data);
  } catch (error) {
    return NextResponse.error();
  }
}

export async function POST(request: NextRequest) {
  if (!request.url) {
    return NextResponse.error();
  }
  let url = request.url.replace("localhost:3000", "localhost:8080");
  let body;
  try {
    body = await request.json();
  } catch (error) {
    return NextResponse.error();
  }
  const config = {
    method: request.method.toLowerCase(),
    url: url,
    data: body,
  };
  try {
    const res = await axios(config);
    if (!res.headers.authorization) {
      return NextResponse.json({ status: 401, message: "Unauthorized" });
    }
    const authToken = res.headers.authorization.split(" ")[1];
    const bffRes = NextResponse.json(res.data);
    bffRes.cookies.set("session_token", authToken, {
      httpOnly: true,
      path: "/",
      expires: new Date(Date.now() + 1000 * 60 * 30),
      sameSite: "lax",
    });
    return bffRes;
  } catch (error: any) {
    if (error.response && error.response.status === 401) {
      return NextResponse.json({ status: 401, message: error.response.data });
    }
    return NextResponse.error();
  }
}
