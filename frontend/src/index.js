import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {BrowserRouter, Route, Routes,} from "react-router-dom";
import RecipesPage from "./recipesPage";
import RecipesElementsPage from "./recipesElementsPage";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<RecipesPage />} />
            <Route path="/recipe/:id" element={<RecipesElementsPage />} />
        </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
