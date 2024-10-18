"use client"; // Ensures the Navbar can handle client-side actions if needed (though it's kept simple here)

import React from "react";
import Link from "next/link";
import Access from "../molecules/access";

export default function Navbar() {
  return (
    <nav className="bg-blue-600 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className="flex items-center space-x-4">
          <Link href="/" passHref className="text-white font-bold text-xl">
            MyApp
          </Link>

          <ul className="flex space-x-4">
            <li>
              <Link
                href="/"
                passHref
                className="text-white hover:text-blue-200"
              >
                Home
              </Link>
            </li>
            <li>
              <Link
                href="/about"
                passHref
                className="text-white hover:text-blue-200"
              >
                About
              </Link>
            </li>
            <li>
              <Link
                href="/contact"
                passHref
                className="text-white hover:text-blue-200"
              >
                Contact
              </Link>
            </li>
          </ul>
        </div>
        <div>
          <Access /> {/* The Access component for Sign In and Sign Up */}
        </div>
      </div>
    </nav>
  );
}
