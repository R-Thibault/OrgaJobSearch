// frontend/middleware.ts
import axios from "axios";
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  // Define a list of public routes (add paths as needed)
  const publicRoutes = ["/login", "/sign-up", "/about", "/contact"];

  // Check if the current path matches any of the public routes
  const isPublicRoute = publicRoutes.some((route) =>
    pathname.startsWith(route)
  );

  // Allow access if the route is public
  if (isPublicRoute) {
    return NextResponse.next();
  }

  // Make a request to the Go backend's /me endpoint to validate the session
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/me`,
      {
        headers: {
          Cookie: req.headers.get("cookie") || "", // Pass all cookies from the request
        },
        withCredentials: true, // Ensures cookies are sent with the request
      }
    );

    // If the response status is not 200, redirect to login
    if (response.status !== 200) {
      return NextResponse.redirect(new URL("/login", req.url));
    }

    // Allow the request to proceed if authenticated
    return NextResponse.next();
  } catch {
    // If there is an error in the fetch request, redirect to login
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

// Apply the middleware only to relevant routes
export const config = {
  matcher: [
    // Apply to all pages, but exclude static files and API routes
    "/((?!_next/static|_next/image|favicon.ico|api).*)",
  ],
};
