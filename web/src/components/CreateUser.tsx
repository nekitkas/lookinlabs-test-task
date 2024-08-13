import React, {useState} from "react";
import { UserRequest } from "../types/user";
import { usePostData } from "../hooks/usePostData";

const url = 'http://localhost:8080/api/v1/users';

export const CreateUser: React.FC = () => {
    const [formData, setFormData] = useState<UserRequest>({ name: '', email: ''})
    const { data, loading, error, postData } = usePostData<UserRequest, any>(
        url,
        { body: formData }
    );

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        postData();
    }

    return (
        <div className="max-w-md mx-auto bg-white shadow-md rounded p-6">
            <h1 className="text-2xl font-bold mb-4">Create User</h1>
            <form onSubmit={handleSubmit}>
                <div className="mb-4">
                    <label htmlFor="name" className="block text-sm font-medium text-gray-700">Name:</label>
                    <input
                        type="text"
                        placeholder="Name"
                        value={formData.name}
                        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>
                <div className="mb-4">
                    <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email:</label>
                    <input
                        type="email"
                        placeholder="Email"
                        value={formData.email}
                        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>
                <button
                    type="submit"
                    disabled={loading}
                    className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white ${loading ? 'bg-gray-500' : 'bg-indigo-600 hover:bg-indigo-700'} focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500`}>
                    {loading ? 'Creating...' : 'Create'}
                </button>
            </form>
            {error && <div className="mt-4 text-red-500">{error}</div>}
            {data && <div className="mt-4 text-green-500">Success! {data.message}</div>}
        </div>
    );
};