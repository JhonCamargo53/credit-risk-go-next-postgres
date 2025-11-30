'use client';
import LoadingPage from "@/components/common/loading/LoadingPage";
import {useRouter } from "next/navigation";

export default function Home() {

  const router = useRouter();

  router.replace("/login");

  return <LoadingPage loadingText='Redirigiendo' />

}
