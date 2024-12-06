-- name: SearchForRecipes :many
SELECT
    *
from
    recipes
where
    title ILIKE '%' || sqlc.arg(search)::text || '%'
    or description ILIKE '%' || sqlc.arg(search)::text || '%';

-- name: GetAllRecipes :many
SELECT
    *
from
    recipes;
