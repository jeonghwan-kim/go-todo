/*global NodeList */
(function (window) {
	'use strict';

	// Get element(s) by CSS selector:
	window.qs = function (selector, scope) {
		return (scope || document).querySelector(selector);
	};
	window.qsa = function (selector, scope) {
		return (scope || document).querySelectorAll(selector);
	};

	// addEventListener wrapper:
	window.$on = function (target, type, callback, useCapture) {
		target.addEventListener(type, callback, !!useCapture);
	};

	// Attach a handler to event for all elements that match the selector,
	// now or in the future, based on a root element
	window.$delegate = function (target, selector, type, handler) {
		function dispatchEvent(event) {
			var targetElement = event.target;
			var potentialElements = window.qsa(selector, target);
			var hasMatch = Array.prototype.indexOf.call(potentialElements, targetElement) >= 0;

			if (hasMatch) {
				handler.call(targetElement, event);
			}
		}

		// https://developer.mozilla.org/en-US/docs/Web/Events/blur
		var useCapture = type === 'blur' || type === 'focus';

		window.$on(target, type, dispatchEvent, useCapture);
	};

	// Find the element's parent with the given tag name:
	// $parent(qs('a'), 'div');
	window.$parent = function (element, tagName) {
		if (!element.parentNode) {
			return;
		}
		if (element.parentNode.tagName.toLowerCase() === tagName.toLowerCase()) {
			return element.parentNode;
		}
		return window.$parent(element.parentNode, tagName);
	};

	// Allow for looping on nodes by chaining:
	// qsa('.foo').forEach(function () {})
  NodeList.prototype.forEach = Array.prototype.forEach;

  window.$http = function (path, method, data, callback) {
    if (!path) {
      throw Error('path is required');
    }
    if (typeof callback !== 'function') {
      throw Error('callback should be function');
    }

    method = method || 'get';
    data = data || null

    var req = new XMLHttpRequest();
    req.open(method, path, true);
    req.onreadystatechange = () => {
      if (req.readyState === 4) {
        if (req.status === 200) {
          try {
            req.data = JSON.parse(req.responseText);
            callback(null, req);
          } catch (err) {
            callback(Error('$http response parse error'));
          }
        } else {
          callback(Error('$http request error'));
        }
      }
    }
    req.send(JSON.stringify(data))
  }
})(window);
