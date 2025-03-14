-- name: GetExpressionByUUID :one
SELECT * FROM expressions
WHERE uuid = $1;


-- name: GetExpressionByID :one
SELECT * FROM expressions
WHERE id = $1;


-- name: FilterExpressions :many
SELECT * FROM expressions
WHERE (
    CASE
        WHEN @expression_filter_op::text = '=' THEN expression = @expression::VARCHAR
        WHEN @expression_filter_op::text = 'in' THEN expression = ANY(@expression_arr::VARCHAR[])
        WHEN @expression_filter_op::text = 'not in' THEN expression <> ALL(@expression_arr::VARCHAR[])
        ELSE TRUE
    END
)
AND (
    CASE
        WHEN @result_filter_op::text = '=' THEN result = @result::REAL
        WHEN @result_filter_op::text = 'in' THEN result = ANY(@result_arr::REAL[])
        WHEN @result_filter_op::text = 'not in' THEN result <> ALL(@result_arr::REAL[])
        ELSE TRUE
    END
)
AND (
    CASE
        WHEN @num_operands_filter_op::text = '=' THEN num_operands = @num_operands::INT
        WHEN @num_operands_filter_op::text = 'in' THEN num_operands = ANY(@num_operands_arr::INT[])
        WHEN @num_operands_filter_op::text = 'not in' THEN num_operands <> ALL(@num_operands_arr::INT[])
        ELSE TRUE
    END
)
AND (
    CASE
        WHEN @difficulty_filter_op::text = '=' THEN difficulty = @difficulty::INT
        WHEN @difficulty_filter_op::text = 'in' THEN difficulty = ANY(@difficulty_arr::INT[])
        WHEN @difficulty_filter_op::text = 'not in' THEN difficulty <> ALL(@difficulty_arr::INT[])
        ELSE TRUE
    END
)
AND (
    CASE
        WHEN @id_filter_op::text = '=' THEN id = @id::INT
        WHEN @id_filter_op::text = 'in' THEN id = ANY(@id_arr::INT[])
        WHEN @id_filter_op::text = 'not in' THEN id <> ALL(@id_arr::INT[])
        ELSE TRUE
    END
)
 AND (
    CASE
        WHEN @uuid_filter_op::text = '=' THEN uuid = @uuid::UUID
        WHEN @uuid_filter_op::text = 'in' THEN uuid = ANY(@uuid_arr::UUID[])
        WHEN @uuid_filter_op::text = 'not in' THEN uuid <> ALL(@uuid_arr::UUID[])
        ELSE TRUE
    END
)
ORDER BY
    CASE WHEN @order_by::text = 'id' AND @order_direction::text = 'asc' THEN id END ASC,
    CASE WHEN @order_by::text = 'id' AND @order_direction::text = 'desc' THEN id END DESC,
    CASE WHEN @order_by::text = 'created_at' AND @order_direction::text = 'asc' THEN created_at END ASC,
    CASE WHEN @order_by::text = 'created_at' AND @order_direction::text = 'desc' THEN created_at END DESC,
    CASE WHEN @order_by::text = 'updated_at' AND @order_direction::text = 'asc' THEN updated_at END ASC,
    CASE WHEN @order_by::text = 'updated_at' AND @order_direction::text = 'desc' THEN updated_at END DESC,
    CASE WHEN @order_by::text = 'deleted_at' AND @order_direction::text = 'asc' THEN deleted_at END ASC,
    CASE WHEN @order_by::text = 'deleted_at' AND @order_direction::text = 'desc' THEN deleted_at END DESC
LIMIT sqlc.narg('limit')::INT
OFFSET sqlc.narg('offset')::INT;
