// frontend/middleware.ts
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
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/me`, {
      headers: {
        Cookie: req.headers.get("cookie") || "", // Pass all cookies from the request
      },
      credentials: "include", // Ensures cookies are sent with the request
    });

    if (response.status === 401) {
      // Redirect to login if unauthenticated
      return NextResponse.redirect(new URL("/login", req.url));
    } else if (response.status === 403) {
      // Redirect to an unauthorized page if user lacks permissions
      return NextResponse.redirect(new URL("/unauthorized", req.url));
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
