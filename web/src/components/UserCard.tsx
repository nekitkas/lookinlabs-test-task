import React from 'react';
import { useFetchData } from "../hooks/useFetchData";
import { User } from "../types/user";

const url = 'http://localhost:8080/api/v1/users';

export const UserCard: React.FC = () => {
    const { data, loading, error } = useFetchData<User[]>(url);

    if (loading) return <div className="text-center">Loading...</div>;
    if (error) return <div className="text-red-500">{error}</div>;

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {data?.map((user: User) => (
                <div key={user.id} className="bg-white shadow-md rounded-lg p-6">
                    <p className="text-lg font-medium text-gray-900">{user.name}</p>
                    <p className="text-gray-500">{user.email}</p>
                </div>
            ))}
        </div>
    );
}