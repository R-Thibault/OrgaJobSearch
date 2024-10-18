"use client";
import Link from "next/link";
import React from "react";

export default function Access() {
  return (
    <div className="flex space-x-4">
      {/* Use a link for Sign In navigation */}
      <Link
        href="/sign-in"
        passHref
        className="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-700"
      >
        Sign In
      </Link>

      {/* Sign Up button can also use a link if it goes to another page */}
      <Link
        href="/sign-up"
        passHref
        className="bg-green-500 text-white font-bold py-2 px-4 rounded hover:bg-green-700"
      >
        Sign Up
      </Link>
    </div>
  );
}
