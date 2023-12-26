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
    data: body,
  };
  try {
    const res = await axios(config);
    return NextResponse.json(res.data);
  } catch (error: any) {
    if (error.response && error.response.status === 400) {
      return NextResponse.json({ status: 400, message: error.response.data });
    }
    return NextResponse.error();
  }
}
