import RecipeItem from "./RecipeItem";
import RecipeItemButton from "./RecipeItemButton";
import { useEffect, useState } from "react";
import axios from "axios";

const RecipesPage = () => {
    const [recipes, setRecipes] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

        const axiosInstance = axios.create({
            baseURL: 'http://localhost:8080/recipes', // Replace with your backend URL and port
            timeout: 1000,
            headers: {
                'Content-Type': 'application/json',
            },
        });

        axiosInstance.get('/')
            .then(response => {
                setRecipes(response.data.data);
                setLoading(false);
            })
            .catch(error => {
                setError(error);
                setLoading(false);
            });
        console.log(recipes);




    return (
        <div>
            <div className="container">
                {recipes.map((recipe) => (
                    <RecipeItem key={recipe.id} id={recipe.id} name={recipe.name} />
                ))}
                <RecipeItemButton />
            </div>
        </div>
    );
};

export default RecipesPage;
