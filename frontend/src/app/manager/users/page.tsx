'use client';

import UserManager from "@/components/modules/user/UserManager";
import { UserContext, UserProvider } from "@/context/UserContext";
import ProtectedRoute from "@/hoc/ProtectedRoute";

const UserPage = () => {

     return <UserProvider><UserManager/></UserProvider>
}

export default ProtectedRoute(UserPage)

