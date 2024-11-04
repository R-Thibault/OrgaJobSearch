"use client";

import Link from "next/link";

export default function BackOfficeNavbar() {
  return (
    <nav className="fixed top-0 left-0 h-full w-64 bg-gray-800 text-white p-4">
      <h2 className="text-xl font-semibold mb-6">Dashboard Navigation</h2>
      <ul>
        <li className="mb-4">Dashboard Home</li>
        <li className="mb-4">User Management</li>
        <li className="mb-4">Settings</li>
        <li className="mb-4">
          <Link href="/my-profile" passHref>
            My Profile
          </Link>
        </li>
      </ul>
    </nav>
  );
}
