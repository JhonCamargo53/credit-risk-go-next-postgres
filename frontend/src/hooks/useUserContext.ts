import { useContext } from 'react';
import { UserContext, UserContextType } from '@/context/UserContext';

export const useUser = (): UserContextType => {
    const context = useContext(UserContext);
    if (context === undefined) {
        throw new Error('useUser debe usarse dentro de un UserProvider');
    }
    return context;
};