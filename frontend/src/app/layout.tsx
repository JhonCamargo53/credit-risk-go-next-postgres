import type { Metadata } from "next";
import "../styles/globals.css";
import { AuthProvider } from "@/context/AuthContext";
import { Toaster } from 'react-hot-toast';
import LayoutWithAuth from "./layout-with-auth";
import { ModalProvider } from "@/context/ModalContext";

export const metadata: Metadata = {
  title: "Product Manager",
  description: "Gestor de rifas",
};


export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          <ModalProvider>
            <LayoutWithAuth>
              {children}
            </LayoutWithAuth>
            <Toaster />
          </ModalProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
