import React, { useState } from 'react';
import { CreateUser } from "./CreateUser";
import { UserCard } from "./UserCard";

export const Home: React.FC = () => {
    const [activeTab, setActiveTab] = useState<'createUser' | 'userCards'>('createUser');

    return (
        <div className="max-w-4xl mx-auto p-6">
            <nav className="mb-4 flex justify-center space-x-4">
                <button
                    onClick={() => setActiveTab('createUser')}
                    className={`px-4 py-2 rounded-md text-white ${activeTab === 'createUser' ? 'bg-indigo-600' : 'bg-gray-600 hover:bg-gray-700'}`}>
                    Create User
                </button>
                <button
                    onClick={() => setActiveTab('userCards')}
                    className={`px-4 py-2 rounded-md text-white ${activeTab === 'userCards' ? 'bg-indigo-600' : 'bg-gray-600 hover:bg-gray-700'}`}>
                    User Cards
                </button>
            </nav>
            <div>
                {activeTab === 'createUser' ? <CreateUser /> : <UserCard />}
            </div>
        </div>
    );
}