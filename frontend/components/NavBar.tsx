"use client"

import Link from "next/link"
import Image from "next/image"

export default function Navbar() {
  return (
    <nav className="sticky top-0 z-100 w-full border-b bg-white/100 backdrop-blur supports-[backdrop-filter]:bg-white/100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 items-center justify-between">
          
          {/* Left Side: Logo and Title */}
          <div className="flex items-center gap-3">
            <Link href="/" className="flex items-center gap-2">
              <div>
                <Image 
                    src="/ICMLogo.png" 
                    alt="ICM Logo" 
                    width={90} 
                    height={90}
                    className="object-contain" 
                />
              </div>
              <span className="hidden sm:inline-block font-bold text-xl tracking-tight text-slate-900">
                Israel Conflict Monitor
              </span>
            </Link>
          </div>

          {/* Right Side: Auth Links */}
          <div className="flex items-center gap-6">
            <Link 
              href="/login" 
              className="text-sm font-medium text-muted-foreground hover:text-primary transition-colors"
            >
              Login
            </Link>
            <Link 
              href="/signup" 
              className="text-sm font-medium text-muted-foreground hover:text-primary transition-colors"
            >
              Sign Up
            </Link>
          </div>

        </div>
      </div>
    </nav>
  )
}