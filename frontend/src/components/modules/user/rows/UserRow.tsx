import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import { GenericTableRow } from '@/types/table'
import { User } from '@/types/user'
import React, { useState } from 'react'
import UserDetailsCard from '../cards/UserDetailsCard';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import { FaAddressCard, FaEye, FaPencilAlt } from 'react-icons/fa';
import ToggleButton from '@/components/common/buttons/ToggleButton';
import CreateUpdateUserForm from '../forms/CreateUpdateUserForm';

interface UserRow extends GenericTableRow<User> {

}

const UserRow: React.FC<UserRow> = ({ index = 0, data, pageData = { currentPage: 1, limit: 10, lastPage: -1 } }) => {

    const [user] = useState(data!);

    const [userDetailsModal, setUserDetailsModal] = useState(false);
    const [userUpdateModal, setUserUpdateModal] = useState(false);

    const toggleUser = async () => {

    }

    return (
        <tr key={user.ID} className={`border-b border-gray-300 hover:bg-gray-200 text-center font-medium ${index % 2 !== 0 ? 'bg-primary/8' : ''}`}>
            <td className="px-4 py-2">{index + 1 + (pageData.currentPage - 1) * pageData.limit}</td>
            <td className="px-4 py-2 whitespace-nowrap">{user.name}</td>
            <td className="px-4 py-2">{user.email}</td>
            <td className="px-4 py-2 whitespace-nowrap">{user.roleId == 1 ? 'ADMINISTRADOR' : 'EMPLEADO'}</td>
            <td className="px-4 py-2">
                <div className="flex items-center justify-center h-full">
                    <Button title="Detalles Usuario" onClick={() => setUserDetailsModal(true)}>
                        <FaEye />
                    </Button>
                </div>
            </td>
            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Editar usuario" onClick={() => setUserUpdateModal(true)}>
                        <FaPencilAlt />
                    </Button>
                </div>
            </td>
            {userDetailsModal &&
                <GenericModal content={<UserDetailsCard user={user} />} isOpen={userDetailsModal}
                    setOpen={setUserDetailsModal} size='md'
                />}

            {userUpdateModal &&
                <GenericModal content={<CreateUpdateUserForm user={user} />}
                    isOpen={userUpdateModal} setOpen={setUserUpdateModal} size='2xl' />}

        </tr>
    )
}

export default UserRow