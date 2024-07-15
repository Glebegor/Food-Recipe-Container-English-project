

const RecipeItem = ({ name, id }) => {
    return (
        <a href={"http://localhost:3000/recipe/" + id} className={"recipeItem"}>
            <div className={"recipeItemTitle"}>{name}</div>
        </a>
    )
}

export default RecipeItem;