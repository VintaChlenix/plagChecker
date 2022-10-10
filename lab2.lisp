(defun nth-element (i lst)
  (if (> i 0) 
      (nth-element (- i 1) (cdr lst))
      (first lst)))
                           
