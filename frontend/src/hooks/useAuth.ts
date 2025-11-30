import { AuthContext, AuthContextType } from '@/context/AuthContext';
import { useContext } from 'react';

export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth debe usarse dentro de un AuthProvider');
    }
    return context;
};
