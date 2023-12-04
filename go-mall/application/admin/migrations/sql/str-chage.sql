-- 逗号分隔字符串转成json字符串数组
UPDATE oms_order_return_apply
SET proof_pics = CASE
     WHEN proof_pics = '' THEN '[]'
     ELSE CONCAT('[\"', REPLACE(proof_pics, ',', '\",\"'), '\"]')
END;