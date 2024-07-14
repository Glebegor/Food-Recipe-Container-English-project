import RecipeItem from "./RecipeItem";
import RecipeItemButton from "./RecipeItemButton";

const RecipesPage = () => {
    return (
        <div>
            <div className={"container"}>
                <RecipeItem />
                <RecipeItem />
                <RecipeItem />
                <RecipeItemButton />
            </div>
        </div>
    )
}

export default RecipesPage;
