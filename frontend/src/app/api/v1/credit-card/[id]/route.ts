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
