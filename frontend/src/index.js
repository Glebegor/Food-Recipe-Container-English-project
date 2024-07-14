import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {Route, Routes,} from "react-router-dom";
import RecipesPage from "./recipesPage";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Routes>
        <Route path="/" element={<RecipesPage />} />
    </Routes>
  </React.StrictMode>
);
