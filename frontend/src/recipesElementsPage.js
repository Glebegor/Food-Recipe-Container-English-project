import RecipeItem from "./RecipeItem";
import RecipeItemButton from "./RecipeItemButton";
import { useEffect, useState } from "react";
import axios from "axios";
import RecipeElementItem from "./RecipeElementItem";

const RecipesElementsPage = () => {
    const [recipes, setRecipes] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const axiosInstance = axios.create({
        baseURL: 'http://localhost:8080/recipesElements', // Replace with your backend URL and port
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
        <div className={"recElemPage"}>
            <div className="container">
                {recipes.map((recipe) => (
                    <RecipeElementItem key={recipe.id} id={recipe.id} name={recipe.name} quantity={recipe.quantity}/>
                ))}
                <RecipeItemButton />
            </div>
        </div>
    );
};

export default RecipesElementsPage;
