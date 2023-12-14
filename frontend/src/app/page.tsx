"use client";
import React from "react";
import LoginForm from "@/components/login";
import Link from "next/link";
import axios from "@/api/axiosConfig";
import { useRouter } from "next/navigation";
import { AxiosResponse } from "axios";
import "./page.module.css";

export default function Home() {
  const router = useRouter();

  return (
    <main>
      <section className="w-full py-12 md:py-24 lg:py-32">
        <div className="container px-4 md:px-6">
          <div className="flex flex-col justify-center space-y-4">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl">
              Super-Metro App
            </h2>
            <p className="max-w-[600px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
              Search for train routes, top-up your account, and use your account
              balance for train rides or purchases at affiliated shops.
            </p>
          </div>
          <div className="grid gap-6 lg:grid-cols-2 xl:grid-cols-3 mt-8">
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Search Routes</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Find the best route from one station to another.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Search Now
                </button>
              </div>
            </div>
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Top-Up</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Easily top-up your account for seamless transactions.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Top-Up Now
                </button>
              </div>
            </div>
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">
                  Use for Rides &amp; Shops
                </h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Use your account balance for train rides or at our affiliated
                  shops.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Learn More
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
      <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-100 dark:bg-gray-800">
        <div className="container px-4 md:px-6">
          <div className="flex flex-col justify-center space-y-4">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl">
              Affiliated Shops
            </h2>
            <p className="max-w-[600px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
              Use your Super-Metro account balance at our wide range of
              affiliated shops and enjoy exclusive discounts.
            </p>
          </div>
          <div className="grid gap-6 lg:grid-cols-2 xl:grid-cols-4 mt-8">
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Shop A</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Description about Shop A.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Visit Shop
                </button>
              </div>
            </div>
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Shop B</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Description about Shop B.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Visit Shop
                </button>
              </div>
            </div>
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Shop C</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Description about Shop C.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Visit Shop
                </button>
              </div>
            </div>
            <div
              className="rounded-lg border bg-card text-card-foreground shadow-sm"
              data-v0-t="card"
            >
              <div className="flex flex-col space-y-1.5 p-6">
                <h3 className="text-lg font-semibold">Shop D</h3>
              </div>
              <div className="p-6">
                <p className="text-sm text-gray-500">
                  Description about Shop D.
                </p>
                <button className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 underline-offset-4 hover:underline h-10 py-2 text-blue-500">
                  Visit Shop
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}
