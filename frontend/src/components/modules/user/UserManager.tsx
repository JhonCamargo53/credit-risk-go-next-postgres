import GenericTable from '@/components/common/tables/GenericTable';
import { DataColumn } from '@/types/table';
import React, { useEffect, useState } from 'react'
import UserRow from './rows/UserRow';
import Button from '@/components/common/buttons/Button';
import { FaPlus } from 'react-icons/fa';
import GenericModal from '@/components/common/modals/GenericModal';
import CreateUpdateUserForm from './forms/CreateUpdateUserForm';
import { useUser } from '@/hooks/useUserContext';
import { generateAxiosErrorToast } from '@/utils/toastUtils';

const userColumns: readonly DataColumn[] = [
  { id: 'index', label: 'N°' },
  { id: 'fullname', label: 'Nombre completo' },
  { id: 'email', label: 'Correo Electrónico', minWidth: 150 },
  { id: 'role', label: 'Rango', minWidth: 150 },
  { id: 'details', label: 'Detalles', minWidth: 100 },
  { id: 'edit', label: 'Editar' }
];
const UserManager = () => {

  const { users, loading, fetchUsers, handlePagination, pageData } = useUser()
  const [createUpdateModal, setCreateUpdateModal] = useState(false);

  useEffect(() => {
    loadUsers()
  }, []);

  const loadUsers = async () => {
    try {
      await fetchUsers()
    } catch (error) {
      generateAxiosErrorToast(error, 'Error al cargar los usuarios', 'Inténtalo nuevamente');
    }
  }

  return (<div className="space-y-4">

    <div className="grid grid-cols-12 gap-4">
      <div className="col-span-2">
        <Button iconLeft={<FaPlus />} className="font-bold w-full h-full" onClick={() => setCreateUpdateModal(true)}>Agregar</Button>
      </div>
    </div>

    <div>
      <GenericTable
        columns={userColumns}
        data={users}
        row={<UserRow />}
        handlePagination={handlePagination}
        loading={loading.fetching}
        pageData={pageData}
      />
    </div>

    {createUpdateModal && <GenericModal content={<CreateUpdateUserForm />} isOpen={createUpdateModal} setOpen={setCreateUpdateModal} size="2xl" />}

  </div>

  )
}

export default UserManager