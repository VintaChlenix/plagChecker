(defun queen-moves (k l m n)
  (if (or (= m k) (= n l) (= (abs (- m k)) (abs (- n l)))) T
       (values m l)))
    