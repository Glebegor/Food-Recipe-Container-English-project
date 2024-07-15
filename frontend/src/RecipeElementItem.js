

const RecipeElementItem = ({ name, id , quantity}) => {
    return (
        <a className={"recipeElementItem"}>
            <div className={"recipeElItemWrapper"}>
                <span>{quantity}</span>
                <div className={"recipeElementItemTitle"}>{name}</div>
            </div>
            <button className={"recipeElementItemButton"}>EDIT</button>
        </a>
    )
}

export default RecipeElementItem;