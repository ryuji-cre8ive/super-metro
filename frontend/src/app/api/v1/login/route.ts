import { NextRequest, NextResponse } from "next/server";

import axios, { AxiosResponse } from "axios";

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
    console.error(error);
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
    console.error(error);
    return NextResponse.error();
  }
  console.log("bodyです", body);
  const config = {
    method: request.method.toLowerCase(),
    url: url,
    headers: Object.fromEntries(request.headers.entries()),
    data: body,
  };
  try {
    const res = await axios(config);
    if (!res.headers["set-cookie"]) {
      return NextResponse.error();
    }
    let setCookieValue = res.headers["set-cookie"][0];
    let [cookieFullValue] = setCookieValue.split(";");
    let [cookieName, cookieValue] = cookieFullValue.split("=");

    let BffResponse = NextResponse.json(res.data);
    BffResponse.cookies.set(cookieName, cookieValue);
    return BffResponse;
  } catch (error) {
    return NextResponse.error();
  }
}
