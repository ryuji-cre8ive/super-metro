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
    return bffRes;
  } catch (error) {
    return new NextResponse(null, { status: 401, statusText: "Unauthorized" });
  }
}
