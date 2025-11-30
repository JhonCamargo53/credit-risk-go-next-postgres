'use client';

import { useRouter } from "next/navigation";

const ManagerPage = () => {
     const router = useRouter();
     return router.replace('manager/dashboard')
}

export default ManagerPage

