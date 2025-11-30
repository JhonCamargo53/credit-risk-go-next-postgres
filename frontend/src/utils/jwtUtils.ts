import { DecodedJwt } from "@/types/jwt"
import { User } from "@/types/user"
import { decodeToken } from "react-jwt"

export const getUserFromToken = (token: string) => {
    return decodeToken(token) as User & DecodedJwt
}