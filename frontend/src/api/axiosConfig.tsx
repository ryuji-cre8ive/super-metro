"use client";
import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

const token = window.localStorage.getItem("session_token");

const axiosConfig: AxiosRequestConfig = {
  baseURL: "http://localhost:8080/api/v1",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
  },
};

const axiosInstance: AxiosInstance = axios.create(axiosConfig);

export default axiosInstance;
